set autocommit = 0;

DELIMITER $$
create procedure transfer(in from_user bigint unsigned, in to_user bigint unsigned, in amount DECIMAL(10,2) )
begin

START TRANSACTION;
-- 账户A
set @b := (select a.balance from accounts a where a.id = from_user);

if @b >= amount then
	update accounts set balance = balance - amount where id = from_user;
	update accounts set balance = balance + amount where id = to_user;
	insert into transactions (from_account_id, to_account_id , amount) values (from_user, to_user, amount);
	commit;
else 
	rollback;
end if;

END$$
DELIMITER ;


call transfer(1, 2, 100);

drop procedure if exists transfer;