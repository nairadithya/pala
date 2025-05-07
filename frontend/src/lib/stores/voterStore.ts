import { writable } from 'svelte/store';

interface Voter {
	voter_id: number;
	first_name: string;
	last_name: string;
	date_of_birth: Date;
	contact_number: string;
}

export const voter = writable<Voter | null>(null);
