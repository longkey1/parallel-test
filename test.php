<?php

$name = !empty($argv[1]) ? $argv[1] : "noname";
echo sprintf('%s %s start' . PHP_EOL, date('Y-m-d H:i:s'), $name);
sleep(60);
echo sprintf('%s %s end' . PHP_EOL, date('Y-m-d H:i:s'), $name);
