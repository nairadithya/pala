<script lang="ts">
	import { onMount } from 'svelte';
	import { voter } from '$lib/stores/voterStore.js';
	import { goto } from '$app/navigation';

	console.log(voter);

	let isLoading = true;
	let error: string | null = null;
	let hasVoted = false;
	let parties = [];
	let selectedParty = null;
	let isSubmitting = false;
	let successMessage = '';

	onMount(async () => {
		if (!$voter || !$voter.data || !$voter.data.id) {
			goto('/');
			return;
		}

		// Check if voter has already voted
		try {
			const response = await fetch(`v1/api/voter/${$voter.data.id}/voted`);
			if (!response.ok) {
				throw new Error(`Error: ${response.status} - ${response.statusText}`);
			}
			const data = await response.json();
			hasVoted = data.hasVoted;

			if (!hasVoted) {
				// Fetch parties from the API
				const partiesResponse = await fetch('v1/api/parties');
				if (!partiesResponse.ok) {
					throw new Error(`Error: ${partiesResponse.status} - ${partiesResponse.statusText}`);
				}
				const partiesData = await partiesResponse.json();
				parties = partiesData.data || [];
			}
		} catch (err) {
			console.error('Error loading data:', err);
			error = err.message || 'Failed to load voting information';
		} finally {
			isLoading = false;
		}
	});

	async function submitVote() {
		if (!selectedParty) {
			error = 'Please select a party to vote for';
			return;
		}

		isSubmitting = true;
		error = null;

		try {
			const response = await fetch('v1/api/vote', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					voterId: $voter.data.id,
					partyId: selectedParty
				})
			});

			if (!response.ok) {
				throw new Error(`Error: ${response.status} - ${response.statusText}`);
			}

			successMessage = 'Your vote has been recorded successfully!';
			hasVoted = true;
		} catch (err) {
			console.error('Vote submission failed:', err);
			error = 'Failed to submit your vote. Please try again.';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div
	class="flex min-h-screen items-center justify-center bg-gradient-to-b from-blue-50 to-indigo-100 p-4"
>
	<div class="w-full max-w-2xl">
		<div class="overflow-hidden rounded-xl bg-white shadow-xl">
			<div class="bg-blue-600 p-6 text-center">
				<h1 class="text-2xl font-bold text-white">Voting Ballot</h1>
				<p class="mt-1 text-blue-100">
					{#if $voter && $voter.data}
						Voter: {$voter.data.firstName} {$voter.data.lastName}
					{/if}
				</p>
			</div>

			<div class="space-y-6 p-6">
				{#if isLoading}
					<div class="flex items-center justify-center py-10">
						<svg
							class="h-8 w-8 animate-spin text-blue-600"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
						>
							<circle
								class="opacity-25"
								cx="12"
								cy="12"
								r="10"
								stroke="currentColor"
								stroke-width="4"
							></circle>
							<path
								class="opacity-75"
								fill="currentColor"
								d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
							></path>
						</svg>
						<span class="ml-2 text-gray-600">Loading voting information...</span>
					</div>
				{:else if error}
					<div class="rounded border-l-4 border-red-500 bg-red-50 p-4">
						<p class="text-sm text-red-700">{error}</p>
					</div>
				{:else if hasVoted}
					<div class="rounded border-l-4 border-green-500 bg-green-50 p-6 text-center">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="mx-auto mb-3 h-12 w-12 text-green-500"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M5 13l4 4L19 7"
							/>
						</svg>
						<h2 class="mb-2 text-xl font-bold text-green-700">Thank You For Voting!</h2>
						<p class="text-green-600">
							Your vote has been recorded and is now part of the election.
						</p>
						<p class="mt-4 text-green-600">You have already cast your ballot for this election.</p>
					</div>
				{:else}
					{#if successMessage}
						<div class="mb-6 rounded border-l-4 border-green-500 bg-green-50 p-4">
							<p class="text-sm text-green-700">{successMessage}</p>
						</div>
					{/if}

					<h2 class="mb-4 text-xl font-semibold text-gray-800">Select Your Party</h2>

					{#if parties.length === 0}
						<p class="text-gray-600">No parties available at this time.</p>
					{:else}
						<div class="space-y-3">
							{#each parties as party}
								<label
									class="flex cursor-pointer items-center rounded-lg border p-4 transition-colors duration-200 hover:bg-blue-50 {selectedParty ===
									party.party_id
										? 'border-blue-500 bg-blue-50'
										: 'border-gray-200'}"
								>
									<input
										type="radio"
										name="party"
										value={party.party_id}
										bind:group={selectedParty}
										class="h-5 w-5 text-blue-600 focus:ring-blue-500"
									/>
									<div class="ml-3 flex-grow">
										<span class="block font-medium text-gray-900">{party.party_name}</span>
										<span class="block text-sm text-gray-500">{party.party_description}</span>
									</div>
								</label>
							{/each}
						</div>

						<button
							on:click={submitVote}
							disabled={isSubmitting || selectedParty === null}
							class="mt-6 flex w-full items-center justify-center rounded-lg bg-blue-600 px-4 py-3 font-medium text-white transition duration-200 hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-70"
						>
							{#if isSubmitting}
								<svg
									class="mr-2 -ml-1 h-4 w-4 animate-spin text-white"
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
								>
									<circle
										class="opacity-25"
										cx="12"
										cy="12"
										r="10"
										stroke="currentColor"
										stroke-width="4"
									></circle>
									<path
										class="opacity-75"
										fill="currentColor"
										d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
									></path>
								</svg>
								Submitting Vote...
							{:else}
								Submit My Vote
							{/if}
						</button>
					{/if}
				{/if}

				<div class="mt-6 text-center">
					<a href="/" class="font-medium text-blue-600 hover:text-blue-800">Return to Home</a>
				</div>
			</div>
		</div>
	</div>
</div>
