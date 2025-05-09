import { writable } from 'svelte/store';

interface Voter {
	voter_id: number;
	first_name: string;
	last_name: string;
	date_of_birth: Date;
	contact_number: string;
	has_voted: boolean;
}

export const voter = writable<Voter | null>(null);
