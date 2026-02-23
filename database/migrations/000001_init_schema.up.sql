CREATE TABLE IF NOT EXISTS roadmaps (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    lifetime TEXT NOT NULL,
    roadmap_id UUID NOT NULL REFERENCES roadmaps(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE,
    is_daily BOOLEAN DEFAULT FALSE,
    lifetime TEXT,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE
);