import System.IO
import Control.Monad
import Data.List.Split
import Data.List (sortBy, nubBy, partition, nub)
import Data.Ord (comparing, Down(..))

main = do
  let n_connections = 2*1000
  fich <- openFile "day8.txt" ReadMode
  full_fich <- hGetContents fich
  let list_string = words full_fich
  let list = divide list_string

  let distance_matrix = makeDistanceMatrix list

  let distance_list = take n_connections $ sortList (makeList distance_matrix)

  let connected_points = getConnectedPoints distance_list
 
  let len_matrix = lenMatrix connected_points
  let result = multiplyList (sortBy (comparing Down) len_matrix)

  print result

  hClose fich

divide :: [String] -> [[Double]]
divide = map divideList
  where
    divideList :: String -> [Double]
    divideList elem = map toDouble (splitOn "," elem)

toDouble :: String -> Double
toDouble = read

makeDistanceMatrix :: [[Double]] -> [[(Double, Int, Int)]]
makeDistanceMatrix points = [ [(distance p1 p2, findPointPos points p1, findPointPos points p2)| p2 <- points ] | p1 <- points ]

distance :: [Double] -> [Double] -> Double
distance [x1,y1,z1] [x2,y2,z2] = sqrt ((x1-x2)^2+(y1-y2)^2+(z1-z2)^2)

minDistances :: [[(Double, Int, Int)]] -> [(Double, Int, Int)]
minDistances = map minDistance

minDistance :: [(Double, Int, Int)] -> (Double, Int, Int)
minDistance [x] = x
minDistance (x:xs) = minComp x (minDistance xs)

minComp :: (Double,Int,Int) -> (Double,Int,Int) -> (Double,Int,Int)
minComp tA@(a,_,_) tB@(b,_,_)
    | a == 0.0 = tB
    | b == 0.0 = tA
    | a > b  = tB
    | a < b  = tA
    | a == b = tA

findMinimumConnexions :: Int -> [(Double,Int,Int)] -> [(Double,Int,Int)]
findMinimumConnexions 0 _ = []
findMinimumConnexions n v = minDistance v : findMinimumConnexions (n-1) (updateList v (minDistance v))

updateList :: [(Double,Int,Int)] -> (Double,Int,Int) -> [(Double,Int,Int)]
updateList list touple =
  let pos = findPos list touple
      (start,_:end) = splitAt pos list
  in start ++ (9999999999999,-1,-1) : end

findPointPos :: [[Double]] -> [Double] -> Int
findPointPos list elt = head [index | (index, e) <- zip [0..] list, e == elt]

findPos :: [(Double,Int,Int)] -> (Double,Int,Int) -> Int
findPos list elt = head [index | (index, e) <- zip [0..] list, e == elt]

sortList :: [(Double,Int,Int)] -> [(Double,Int,Int)]
sortList = sortBy minCompList

minCompList :: (Double,Int,Int) -> (Double,Int,Int) -> Ordering
minCompList (a,_,_) (b,_,_)
    | a == 0.0 && b == 0.0 = EQ
    | a == 0.0 = GT
    | b == 0.0 = LT
    | a < b = LT
    | a > b = GT
    | a == b = LT

makeList :: [[(Double,Int,Int)]] -> [(Double,Int,Int)]
makeList = concat

getConnectedPoints :: [(Double, Int, Int)] -> [[Int]]
getConnectedPoints list = foldl mergeGroups [] pairs
  where
    pairs = [(x, y) | (_, x, y) <- list]

mergeGroups :: [[Int]] -> (Int, Int) -> [[Int]]
mergeGroups groups (x, y) =
    let (exists, rest) = partition (\group -> x `elem` group || y `elem` group) groups
        newGroup = nub $ concat exists ++ [x, y]
    in newGroup : rest

lenList :: [Int] -> Int
lenList = foldl (\x _ -> x + 1) 0

lenMatrix :: [[Int]] -> [Int]
lenMatrix = map lenList

multiplyList :: [Int] -> Int
multiplyList (x1:x2:x3:_) = x1 * x2 * x3
