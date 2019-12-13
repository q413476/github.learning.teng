<?php
/**
 * Created by IntelliJ IDEA.
 * User: tengteng
 * Date: 2019-12-10
 * Time: 17:40
 */


$r = $_GET;

echo json_encode([
    [1, 2, 3, 3, 4, 5, 6],
    "name" => "张三",
    $r
]);