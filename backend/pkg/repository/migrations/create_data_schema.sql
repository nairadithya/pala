CREATE TABLE IF NOT EXISTS "party" (
    party_id SERIAL PRIMARY KEY,
    party_name VARCHAR(255) NOT NULL UNIQUE,
    party_description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "voter" (
    voter_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE,
    contact_number VARCHAR(50),
    registration_date DATE DEFAULT CURRENT_DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "candidate" (
    candidate_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    party_id INT,
    bio TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (party_id) REFERENCES party(party_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "election" (
    election_id SERIAL PRIMARY KEY,
    election_name VARCHAR(255) NOT NULL,
    election_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "vote" (
    vote_id SERIAL PRIMARY KEY, 
    election_id INT NOT NULL, 
    voter_id INT NOT NULL, 
    candidate_id INT NOT NULL, 
    voted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (election_id) REFERENCES election(election_id) ON DELETE CASCADE, 
    FOREIGN KEY (voter_id) REFERENCES voter(voter_id) ON DELETE CASCADE, 
    FOREIGN KEY (candidate_id) REFERENCES candidate(candidate_id) ON DELETE CASCADE, 

    UNIQUE (election_id, voter_id)
);
