@echo off

rm -r parser
rm -r java_build

:: 1. Generate Go parser files for your Go application
java -jar antlr-4.13.2-complete.jar -Dlanguage=Go -o parser ANTLR4_Code/ParserSea.g4

:: 3. Generate Java parser files into \java_build
java -jar antlr-4.13.2-complete.jar -o java_build ANTLR4_Code/ParserSea.g4

:: 4. Compile the generated Java files
javac -cp "antlr-4.13.2-complete.jar" java_build/*.java
echo Build complete!