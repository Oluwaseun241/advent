import scala.io.Source
// start from the smallest
// check how many times the left hand side appears on the right
// multiply left hand side by number of occurence and store it
//do this for all the list
// add the total similarity score
@main def main(): Unit = {
  // example
  // val input = List(
  //   (3,4),
  //   (4,3),
  //   (2,5),
  //   (1,3),
  //   (3,9),
  //   (3,3)
  // )

  // read txt file
  val filename = "./input.txt"
  val input = Source.fromFile(filename).getLines().toList.map { line =>
    val parts = line.split("\\s+").map(_.trim.toInt)
    (parts(0), parts(1))
  }

  var lhSide = input.map(_._1)
  var rhSide = input.map(_._2)
  
  // calculate similarity score
  val similarity = lhSide.map { left  =>
    val count = rhSide.count(_ == left)
    
    left * count
  }.sum

  println(s"Similarity score: $similarity")
}
