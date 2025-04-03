<script>
	import { Input, Label, Helper, Button, Checkbox, A } from 'flowbite-svelte';
</script>

<div class="text-primary-500 text-center text-5xl">Welcome To Pala</div>
<div class="my-2 text-2xl">Registration For Elections</div>
<form>
	<div class="mb-6 grid gap-6 md:grid-cols-2">
		<div>
			<Label for="first_name" class="mb-2">First name</Label>
			<Input type="text" id="first_name" placeholder="John" required />
		</div>
		<div>
			<Label for="last_name" class="mb-2">Last name</Label>
			<Input type="text" id="last_name" placeholder="Doe" required />
		</div>
	</div>
	<div class="mb-6">
		<Label for="voter_id" class="mb-2">Voter ID</Label>
		<Input type="email" id="voter_id" placeholder="VXX" required />
	</div>
	<div class="mb-6">
		<Label for="password" class="mb-2">Password</Label>
		<Input type="password" id="password" placeholder="•••••••••" required />
	</div>
	<div class="mb-6">
		<Label for="confirm_password" class="mb-2">Confirm password</Label>
		<Input type="password" id="confirm_password" placeholder="•••••••••" required />
	</div>
	<Checkbox class="mb-6" required>
		<p>
			I agree with the <A href="/" class="text-primary-700 dark:text-primary-600 hover:underline"
				>terms and conditions</A
			>
		</p>
		.
	</Checkbox>
	<Button type="submit">Submit</Button>
</form>
