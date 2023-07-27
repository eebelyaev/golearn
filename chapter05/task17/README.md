Упражнение 5.17. Напишите вариативную функцию ElementsByTagName, которая для данного дерева узла HTML и нуля или нескольких имен возвращает все элементы, которые соответствуют одному из этих имен. Вот два примера вызова такой
функции:
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "hi", "h2", "h3", "h4")