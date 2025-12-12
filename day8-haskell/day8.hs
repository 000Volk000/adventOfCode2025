import System.IO
import Control.Monad
import Data.List.Split

main = do
  let n_connections = 10
  fich <- openFile "day8-example.txt" ReadMode
  full_fich <- hGetContents fich
  let list_string = words full_fich
  let list = divide list_string

  let distance_matrix = makeDistanceMatrix list

  let minimum_distances = minDistances distance_matrix
  print minimum_distances

  print (minDistance minimum_distances)
  hClose fich

divide :: [String] -> [[Double]]
divide = map divideList
  where
    divideList :: String -> [Double]
    divideList elem = map toDouble (splitOn "," elem)

toDouble :: String -> Double
toDouble = read

makeDistanceMatrix :: [[Double]] -> [[Double]]
makeDistanceMatrix points = [ [ distance p1 p2 | p2 <- points ] | p1 <- points ]

distance :: [Double] -> [Double] -> Double
distance [x1,y1,z1] [x2,y2,z2] = sqrt ((x1-x2)^2+(y1-y2)^2+(z1-z2)^2)

minDistances :: [[Double]] -> [Double]
minDistances = map minDistance

minDistance :: [Double] -> Double
minDistance [x] = x
minDistance (x:xs) = minComp x (minDistance xs)

minComp :: Double -> Double -> Double
minComp a b
    | a == 0.0 = b
    | b == 0.0 = a
    | a > b  = b
    | a < b  = a
    | a == b = a
