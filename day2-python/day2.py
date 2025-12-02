if __name__ == "__main__":
    fich = open("day2.txt", "r")

    wrongIds = []
    rangesAll = fich.read()
    rangesSplitted = rangesAll.split(",")

    for rTotal in rangesSplitted:
        r = rTotal.split("-")
        for i in range(int(r[0]), int(r[1]) + 1):
            strI = str(i)
            if len(strI) % 2 == 0:
                startN = strI[: int(len(strI) / 2)]
                endN = strI[int(len(strI) / 2) :]
                if startN == endN:
                    wrongIds.append(i)

    sum = 0
    for el in wrongIds:
        sum += el

    print(f"The total sum is: {sum}")
