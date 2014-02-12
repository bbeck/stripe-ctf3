import AssemblyKeys._

// disable using the Scala version in output paths and artifacts
crossPaths := false

resolvers += "twitter" at "http://maven.twttr.com/"

resolvers += "retronym-releases" at "http://retronym.github.com/repo/releases"

libraryDependencies += "com.twitter" %% "twitter-server" % "1.0.2"

libraryDependencies += "com.twitter" %% "finagle-core" % "6.5.2"

libraryDependencies += "com.twitter" %% "finagle-stream" % "6.5.2"

libraryDependencies += "org.rogach" %% "scallop" % "0.9.4"

libraryDependencies += "com.yammer.dropwizard" % "dropwizard-core" % "0.6.2"

assemblySettings

jarName in assembly := "level3.jar"

outputPath in assembly := file("./level3.jar")

//mergeStrategy in assembly <<= (mergeStrategy in assembly) { (old) =>
//  {
//    case PathList("com", "twitter", "common", "args", xs @ _*) => MergeStrategy.last
//    case PathList(ps @ _*) if ps.last endsWith ".html" => MergeStrategy.first
//    case PathList("META-INF", "MANIFEST.MF") => MergeStrategy.discard
//    case m if m.toLowerCase.matches("meta-inf.*\\.sf$") => MergeStrategy.discard
//    case m if m.toLowerCase.endsWith("manifest.mf") => MergeStrategy.discard
//    case _ => MergeStrategy.first
//  }
//}

mergeStrategy in assembly := {
  case m if m.toLowerCase.endsWith("manifest.mf") => MergeStrategy.discard
  case m if m.toLowerCase.matches("meta-inf.*\\.sf$") => MergeStrategy.discard
  case "log4j.properties" => MergeStrategy.discard
  case m if m.toLowerCase.startsWith("meta-inf/services/") => MergeStrategy.filterDistinctLines
  case "reference.conf" => MergeStrategy.concat
  case _ => MergeStrategy.first
}
