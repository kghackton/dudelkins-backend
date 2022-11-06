CREATE INDEX CONCURRENTLY IF NOT EXISTS  applications_idx
    ON applications(
            unom
        , defect_id
        , entrance
        , floor
        , flat
        , closed_at
        );