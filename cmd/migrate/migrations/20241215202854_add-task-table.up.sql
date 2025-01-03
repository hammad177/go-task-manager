CREATE TABLE IF NOT EXISTS tasks (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `status` ENUM('TODO', 'IN_PROGRESS', 'TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
    `project_id` INT UNSIGNED NOT NULL,
    `assigned_to` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (assigned_to) REFERENCES users(id)
);