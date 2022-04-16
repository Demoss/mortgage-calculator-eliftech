CREATE TABLE banks
(
    name             varchar(50) not null unique,
    rate             int         not null,
    max_loan         int         not null,
    min_down_payment int         not null,
    loan_term        int         not null
);

INSERT INTO banks (name, rate, max_loan, min_down_payment, loan_term)
VALUES ('PrivatBank',18,1000000,20,24),
       ('UKRSIBBANK',15,500000,20,12),
       ('OSCHADBANK',20,1500000,30,24),
       ('PUMB',15,100000,15,18),
       ('RAIFFAISEN',18,2000000,25,24)