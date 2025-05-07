<script lang="ts">
	let firstName: string = '';
	let lastName: string = '';
	let dateOfBirth: string = '';
	let contactNumber: string = '';

	let message: string | null = null;
	let isSubmitting: boolean = false;

	const backendEndpoint: string = '/v1/api/voter/';

	async function handleSubmit() {
		event?.preventDefault();

		isSubmitting = true;
		message = 'Submitting...';

		const date = new Date(dateOfBirth);
		const rfc339 = date.toISOString();

		const dataToSend = {
			first_name: firstName,
			last_name: lastName,
			date_of_birth: rfc339, // Sending 'YYYY-MM-DDTHH:mm'
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
				// --- PLACEHOLDER: Handle successful response ---
				// You might parse the response body if the backend returns data
				// const result = await response.json();
				message = 'Data submitted successfully!';
				// Optionally clear the form:
				// firstName = '';
				// lastName = '';
				// dateOfBirth = '';
				// contactNumber = '';
			} else {
				let errorText = `Error: ${response.status} ${response.statusText}`;
				try {
					const errorBody = await response.json();
					if (errorBody && errorBody.message) {
						errorText = `Error: ${response.status} - ${errorBody.message}`;
					} else if (errorBody) {
						errorText = `Error: ${response.status} - ${JSON.stringify(errorBody)}`;
					}
				} catch (e) {
					// Ignore JSON parsing error if response isn't JSON
				}
				message = errorText;
			}
		} catch (error) {
			// --- PLACEHOLDER: Handle network errors ---
			console.error('Fetch error:', error);
			message = `Network error: ${error instanceof Error ? error.message : String(error)}`;
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div class="container mx-auto max-w-md p-4">
	<h1 class="mb-6 text-center text-2xl font-bold">Voter Registration</h1>

	<form on:submit|preventDefault={handleSubmit} class="space-y-4">
		<div>
			<label for="firstName" class="block text-sm font-medium text-gray-700">First Name</label>
			<input
				type="text"
				id="firstName"
				bind:value={firstName}
				required
				class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 focus:outline-none sm:text-sm"
			/>
		</div>

		<div>
			<label for="lastName" class="block text-sm font-medium text-gray-700">Last Name</label>
			<input
				type="text"
				id="lastName"
				bind:value={lastName}
				required
				class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 focus:outline-none sm:text-sm"
			/>
		</div>

		<div>
			<label for="dateOfBirth" class="block text-sm font-medium text-gray-700">Date of Birth</label>
			<input
				type="date"
				id="dateOfBirth"
				bind:value={dateOfBirth}
				required
				class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 focus:outline-none sm:text-sm"
			/>
			<p class="mt-1 text-xs text-gray-500">Select date and time (YYYY-MM-DDTHH:MM)</p>
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
				class="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 focus:outline-none sm:text-sm"
			/>
		</div>

		<div>
			<button
				type="submit"
				disabled={isSubmitting}
				class="flex w-full justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
			>
				{isSubmitting ? 'Sending...' : 'Submit Voter Details'}
			</button>
		</div>
	</form>

	{#if message}
		<p
			class="mt-4 text-center text-sm {message.startsWith('Error') ||
			message.startsWith('Network error')
				? 'text-red-600'
				: 'text-green-600'}"
		>
			{message}
		</p>
	{/if}
</div>
