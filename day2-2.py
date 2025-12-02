import textwrap

if __name__ == "__main__":
    fich = open("day2.txt", "r")

    wrongIds = []
    rangesAll = fich.read()
    rangesSplitted = rangesAll.split(",")

    for rTotal in rangesSplitted:
        r = rTotal.split("-")
        for i in range(int(r[0]), int(r[1]) + 1):
            strI = str(i)
            for s in range(1, len(strI)):
                if i not in wrongIds:
                    equals = textwrap.wrap(strI, s)
                    cont = 0
                    start = int(equals[0])
                    index = 1
                    while index < len(equals):
                        if equals[0] == equals[index]:
                            cont += 1
                        index += 1
                    if cont == (len(equals) - 1):
                        wrongIds.append(i)

    sum = 0
    for el in wrongIds:
        sum += el

    print(f"The total sum is: {sum}")
