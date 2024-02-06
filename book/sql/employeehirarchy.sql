-- -------------------------------------------------------------
-- Get all customers that don't have an order
-- -------------------------------------------------------------

select * from customers c 
left join orders o  on c.customer_id  = o.customer_id 
where o.customer_id is null

-- -------------------------------------------------------------
-- Get all the orders (from orders table) that contain ‘Socks’
-- -------------------------------------------------------------
select o.order_id , od.product_id , p.product_name , p.unit_price  from orders o 
inner join order_details od on o.order_id = od.order_id 
inner  join products p on p.product_id = od.product_id 
where p.product_name ilike 'Socks'

-- -------------------------------------------------------------
-- Get the customer names with the most orders
-- -------------------------------------------------------------
select c.customer_name , count(1)  from customers c
inner join orders o on c.customer_id = o.customer_id
group by c.customer_id
order by count(1) desc


-- -------------------------------------------------------------
-- Get the customer names with the most orders only
-- -------------------------------------------------------------

with max_order_cte AS
(
	select count(1) from orders o 
	group by customer_id
	order by count(1) desc 
	limit 1
)

select c.customer_id ,c.customer_name, COUNT(o.order_id) from customers c inner join orders o
on c.customer_id  = o.customer_id
group by c.customer_id, c.customer_name 
having COUNT(o.order_id) = (select * from max_order_cte);

-- Second Solution

select c.customer_id ,c.customer_name , count(1) from orders o 
inner join customers c on c.customer_id = o.customer_id 
group by c.customer_name , c.customer_id 
having count(1) = (select count(1) from orders 
					group by customer_id 
					order by count(1) desc limit 1);

-- -------------------------------------------------------------
-- Get the top 3 most ordered products
-- -------------------------------------------------------------

select od.product_id, p.product_name, COUNT(1)  from orders o 
inner join order_details od 
on o.order_id = od.order_id
inner join products p on p.product_id  = od.product_id 
group by od.product_id, p.product_name  
order by COUNT(*) desc limit 3;

-- -------------------------------------------------------------
-- Get the last shipped item (with date and product name)
-- for every customer
-- -------------------------------------------------------------
with cte_order_with_shipdate AS(
	select o.customer_id, MAX(o.shipped_date) as recent_ship_date from orders o 
	group by o.customer_id
)

select c.customer_name, cs.recent_ship_date, p.product_name  from cte_order_with_shipdate cs
inner join orders o on cs.customer_id = o.customer_id and o.shipped_date  = cs.recent_ship_date
inner join order_details od on od.order_id  = o.order_id
inner join products p on p.product_id  = od.product_id
inner join customers c on c.customer_id = cs.customer_id;


-- Second Solution with partition
with recent_orders AS(
	select distinct first_value (customer_id) over (partition by customer_id order by shipped_date desc) as customer_id,
	first_value (order_id) over (partition by customer_id order by shipped_date desc) as order_id,
	first_value (shipped_date) over (partition by customer_id order by shipped_date desc) as shipped_date
	
	from orders
)

select c.customer_name , ro.shipped_date, p.product_name  from recent_orders ro
join order_details od on ro.order_id = od.order_id 
join products p on p.product_id =od.product_id 
join customers c on c.customer_id = ro.customer_id;


-- -------------------------------------------------------------
-- Get TOP 3 product qualities ordered
-- -------------------------------------------------------------
select product_name, SUM(quantity), rank() over(order by sum(quantity) desc) as my_rank from order_details od
inner join products p on od.product_id = p.product_id
group by product_name
order by SUM(quantity) desc
limit 3;


-- Another Solution
select p.product_name, total_quantity, my_rank
from(
	select product_id, sum(quantity) total_quantity, rank() over(order by sum(quantity) desc) as my_rank
	from order_details od 
	group by product_id
) a
join products p on p.product_id = a.product_id
where my_rank<=3
order by my_rank;