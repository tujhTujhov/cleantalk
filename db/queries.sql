-- Для IPv4
SELECT *
FROM ipv4_networks
WHERE INET_ATON('some_ip') & (4294967295 << (32 - subnet_mask)) = network
ORDER BY subnet_mask DESC LIMIT 1;

-- Для IPv6
SELECT *
FROM ipv6_networks
WHERE INET6_ATON('some_ip') & (CONV(REPEAT('1', subnet_mask), 2, 10)) = network
ORDER BY subnet_mask DESC LIMIT 1;