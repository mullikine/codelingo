<?php
class MyClass
{
    const lcconstant = 'constant value';
    const CONSTANT = 'constant value';

    function showConstant() {
        echo  self::CONSTANT . "\n";
        echo  self::lcconstant . "\n";
    }
}

echo MyClass::CONSTANT . "\n";
echo MyClass::lcconstant . "\n";

$classname = "MyClass";
echo $classname::CONSTANT . "\n"; // As of PHP 5.3.0
echo $classname::lcconstant . "\n"; // As of PHP 5.3.0

$class = new MyClass();
$class->showConstant();

echo $class::CONSTANT."\n"; // As of PHP 5.3.0
echo $class::lcconstant."\n"; // As of PHP 5.3.0

class foo {
    // As of PHP 5.3.0
    const BAR = <<<'EOT'
bar
EOT;
    const lcbar = <<<'EOT'
bar
EOT;
    // As of PHP 5.3.0
    const BAZ = <<<EOT
baz
EOT;
    // As of PHP 5.3.0
    const lcbaz = <<<EOT
baz
EOT;
}
?>