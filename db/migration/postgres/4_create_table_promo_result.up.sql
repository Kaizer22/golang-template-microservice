CREATE TABLE IF NOT EXISTS promotion_result (
    promo_id INT,
    winner_id INT,
    prize_id INT,
    PRIMARY KEY (promo_id, winner_id, prize_id),
    CONSTRAINT fk_promo FOREIGN KEY (promo_id) REFERENCES promotion(id),
    CONSTRAINT fk_winner FOREIGN KEY (winner_id) REFERENCES participant(id),
    CONSTRAINT fk_prize FOREIGN KEY (prize_id) REFERENCES prize(id)
);
