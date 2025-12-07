Explored = {}

function Down_the_matrix(lines, i, index, switch)
	if i >= #lines then
		Explored[i .. "," .. index] = 0
		return 0
	end

	if Explored[i .. "," .. index] then
		return Explored[i .. "," .. index]
	end

	if switch[string.sub(lines[i], index, index)] then
		Explored[i .. "," .. index] = 1
			+ Down_the_matrix(lines, i + 1, index - 1, switch)
			+ Down_the_matrix(lines, i + 1, index + 1, switch)
	else
		Explored[i .. "," .. index] = Down_the_matrix(lines, i + 1, index, switch)
	end

	return Explored[i .. "," .. index]
end

local fich = io.open("day7.txt", "r")
if fich == nil then
	error("Error opening the file")
end

local lines = {}
local i = 1

for line in fich:lines() do
	lines[i] = line
	i = i + 1
end

fich:close()

local index = string.find(lines[1], "S")
if index == nil then
	error("Couldnt find the S on the first line")
end
table.remove(lines, 1)

local switch = {
	["."] = false,
	["^"] = true,
}
local cont = Down_the_matrix(lines, 1, index, switch) + 1

print("The total number of divisions is: " .. cont)
