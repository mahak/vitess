drop table if exists onlineddl_test;
create table onlineddl_test (
  id int auto_increment,
  c1 int not null,
  c2 int not null,
  primary key (id)
) auto_increment=1;

drop event if exists onlineddl_test;
delimiter ;;
create event onlineddl_test
  on schedule every 1 second
  starts current_timestamp
  ends current_timestamp + interval 60 second
  on completion not preserve
  enable
  do
begin
  insert into onlineddl_test values (null, 11, 23);
  insert into onlineddl_test values (null, 13, 23);
  insert into onlineddl_test values (null, floor(rand()*pow(2,32)), floor(rand()*pow(2,32)));
end ;;
