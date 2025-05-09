<script lang="ts">
	let firstName: string = '';
	let lastName: string = '';
	let dateOfBirth: string = '';
	let contactNumber: string = '';

	let message: string | null = null;
	let createdVoterId: number | null = null;
	let isSubmitting: boolean = false;

	const backendEndpoint: string = '/v1/api/voter/';

	async function handleSubmit() {
		event?.preventDefault();

		isSubmitting = true;
		message = 'Submitting...';

		const date = new Date(dateOfBirth);
		const rfc3339 = date.toISOString();

		const dataToSend = {
			first_name: firstName,
			last_name: lastName,
			date_of_birth: rfc3339,
			contact_number: contactNumber
		};
		console.log(dataToSend);
		try {
			const response = await fetch(backendEndpoint, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(dataToSend)
			});

			if (response.ok) {
				const result = await response.json();
				createdVoterId = result.voter_id;
				message = `Data submitted successfully! Voter ID: ${createdVoterId}`;
				firstName = '';
				lastName = '';
				dateOfBirth = '';
				contactNumber = '';
			} else {
				let errorText = `Error: ${response.status} ${response.statusText}`;
				try {
					const errorBody = await response.json();
					if (errorBody && errorBody.message) {
						errorText = `Error: ${response.status} - ${errorBody.message}`;
					} else if (errorBody) {
						errorText = `Error: ${response.status} - ${JSON.stringify(errorBody)}`;
					}
				} catch (e) {}
				message = errorText;
				createdVoterId = null;
			}
		} catch (error) {
			createdVoterId = null;
			console.error('Fetch error:', error);
			message = `Network error: ${error instanceof Error ? error.message : String(error)}`;
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div
	class="flex min-h-screen items-center justify-center bg-gradient-to-b from-blue-50 to-indigo-100 p-4"
>
	<div class="w-full max-w-md">
		<div class="overflow-hidden rounded-xl bg-white shadow-xl">
			<div class="bg-blue-600 p-6 text-center">
				<h1 class="text-2xl font-bold text-white">Voter Registration</h1>
			</div>
			<form on:submit|preventDefault={handleSubmit} class="space-y-6 p-6">
				<div>
					<label for="firstName" class="block text-sm font-medium text-gray-700">First Name</label>
					<input
						type="text"
						id="firstName"
						bind:value={firstName}
						required
						class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-blue-500 focus:outline-none sm:text-sm"
					/>
				</div>
				<div>
					<label for="lastName" class="block text-sm font-medium text-gray-700">Last Name</label>
					<input
						type="text"
						id="lastName"
						bind:value={lastName}
						required
						class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-blue-500 focus:outline-none sm:text-sm"
					/>
				</div>
				<div>
					<label for="dateOfBirth" class="block text-sm font-medium text-gray-700"
						>Date of Birth</label
					>
					<input
						type="date"
						id="dateOfBirth"
						bind:value={dateOfBirth}
						required
						class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-blue-500 focus:outline-none sm:text-sm"
					/>
					<p class="mt-1 text-xs text-gray-500">Enter your date of birth (YYYY-MM-DD)</p>
				</div>
				<div>
					<label for="contactNumber" class="block text-sm font-medium text-gray-700"
						>Contact Number</label
					>
					<input
						type="text"
						id="contactNumber"
						bind:value={contactNumber}
						required
						class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-blue-500 focus:ring-blue-500 focus:outline-none sm:text-sm"
					/>
				</div>
				<div>
					<button
						type="submit"
						disabled={isSubmitting}
						class="flex w-full justify-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
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
							                            Submitting...
						{:else}
							Submit Voter Details
						{/if}
					</button>
				</div>

				{#if message}
					<p
						class="mt-4 text-center text-sm {message.startsWith('Error') ||
						message.startsWith('Network error')
							? 'text-red-600'
							: 'text-green-600'}"
					>
						{message}
						<a href="/" class="block underline">Proceed back to login</a>
					</p>
				{/if}
			</form>
		</div>
	</div>
</div>
