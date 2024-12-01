import scala.io.Source
// start from the smallest
// compare left hand side to right
// and get the distance between them and store it
//do this for all the list
// add the total distance
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

  var lhSide = input.map(_._1).sorted
  var rhSide = input.map(_._2).sorted

  val distance = lhSide.zip(rhSide).map { case (left, right) =>
    math.abs(left - right)
  }

  val totalDistance = distance.sum
  println(s"Total distance: $totalDistance")
}
