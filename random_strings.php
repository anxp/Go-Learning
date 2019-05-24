<?php
/**
 * Created by PhpStorm.
 * User: andrey
 * Date: 2019-05-23
 * Time: 22:37
 */

/**
 * Generate a random string, using a cryptographically secure
 * pseudorandom number generator (random_int)
 *
 * For PHP 7, random_int is a PHP core function
 * For PHP 5.x, depends on https://github.com/paragonie/random_compat
 *
 * @param int $length      How many characters do we want?
 * @param string $keyspace A string of all possible characters to select from
 * @return bool|string
 */
function randString(int $length, $keyspace = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'): string {
    $pieces = [];
    $max = mb_strlen($keyspace, '8bit') - 1;
    try {
        for ($i = 0; $i < $length; ++$i) {
            $pieces [] = $keyspace[random_int(0, $max)];
        }
    } catch (Exception $e) {
        return FALSE;
    }
    return implode('', $pieces);
}

function printMemory(): string {
    /* Currently used memory */
    $memUsage = memory_get_usage();

    /* Peak memory usage */
    $memPeak = memory_get_peak_usage();

    return 'Current memory using: ' . round($memUsage / 1024) . 'KB; Peak usage: ' . round($memPeak / 1024) . 'KB;';
}

$hashesPool = [];
$hashesCount = 1000000;

echo 'Before big array created: '. printMemory() . PHP_EOL;

$t1 = time();
for ($i=0; $i < $hashesCount; $i++) {
    $hashesPool[$i] = randString(10);
}
$t2 = time();

echo 'After big array created: '. printMemory() . PHP_EOL;

unset($hashesPool);

echo "Execution time is: " . ($t2 - $t1) . "sec.";