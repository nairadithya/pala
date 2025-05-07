<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { voter } from '$lib/stores/voterStore.js';

	console.log(voter);

	let voterId = '';
	let isLoading = false;
	let error: string | null = null;

	async function lookupVoter() {
		if (!voterId.trim()) {
			error = 'Please enter a valid Voter ID';
			return;
		}

		isLoading = true;
		error = null;

		try {
			const response = await fetch(`/api/GetVoters/`);

			if (!response.ok) {
				throw new Error(`Error: ${response.status} - ${response.statusText}`);
			}

			const voterData = await response.json();
			voter.set(voterData);

			goto('/vote');
		} catch (err) {
			console.error('Voter lookup failed:', err);
			error = 'Could not find voter with that ID. Please check and try again.';
		} finally {
			isLoading = false;
		}
	}
</script>

<div
	class="flex min-h-screen items-center justify-center bg-gradient-to-b from-blue-50 to-indigo-100 p-4"
>
	<div class="w-full max-w-md">
		<div class="overflow-hidden rounded-xl bg-white shadow-xl">
			<div class="bg-blue-600 p-6 text-center">
				<h1 class="text-2xl font-bold text-white">Voter Verification</h1>
				<p class="mt-1 text-blue-100">Enter your Voter ID to proceed</p>
			</div>

			<form on:submit|preventDefault={lookupVoter} class="space-y-4 p-6">
				<div>
					<label for="voterId" class="mb-1 block text-sm font-medium text-gray-700">Voter ID</label>
					<input
						id="voterId"
						type="text"
						bind:value={voterId}
						class="w-full rounded-lg border border-gray-300 px-4 py-3 transition duration-200 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
						placeholder="Enter your Voter ID number"
						autocomplete="off"
					/>
				</div>

				{#if error}
					<div class="rounded border-l-4 border-red-500 bg-red-50 p-4">
						<p class="text-sm text-red-700">{error}</p>
					</div>
				{/if}

				<button
					type="submit"
					class="flex w-full items-center justify-center rounded-lg bg-blue-600 px-4 py-3 font-medium text-white transition duration-200 hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:outline-none"
					disabled={isLoading}
				>
					{#if isLoading}
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
						Verifying...
					{:else}
						Continue to Voting
					{/if}
				</button>
			</form>

			<div class="border-t border-gray-200 bg-gray-50 px-6 py-4"></div>
		</div>
	</div>
</div>
