Упражнение 6.5. Типом каждого слова, используемого в IntSet, является uint64, но 64-разрядная арифметика может быть неэффективной на 32-разрядных платформах. Измените программу так, чтобы она использовала тип uint, который представляет собой наиболее эффективный беззнаковый целочисленный тип для данной платформы. Вместо деления на 64 определите константу, в которой хранится эффективный размер uint в битах, 32 или 64. Для этого можно воспользоваться, возможно, слишком умным выражением 32<<(^uint(0)>>63).