<?php
require "vendor/autoload.php";
header("Access-Control-Allow-Origin: *");

use GeoIp2\Database\Reader;

$cityDbReader = new Reader('geoip2country.mmdb');
$cache = new Memcached;

$ip = $_GET['ip'];

if (!filter_var($ip, FILTER_VALIDATE_IP)) {
    echo ("is not a valid ip");
    return;
}

$data = $cache->get($ip);
if (!$data) {
    $data = $cityDbReader->country($ip)->jsonSerialize();
    $res = [
        'country_name' => $data['country']['names']['en'],
        'iso_code' => $data['country']['iso_code'],
        'geoname_id' => $data['country']['geoname_id']
    ];
    $cache->set($ip, $res);
    echo json_encode($res);
} else {
    $res = $cache->get($ip);
    echo json_encode($res);
}


