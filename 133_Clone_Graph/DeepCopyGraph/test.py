

def main():
    digraph = {
        "B": ["E","C"],
        "E":["A"],
        "C":["D"],
        "A":["C","D"],
        "D":[],
    }

    #using list comprehension  to erect/create a initalized Dictionary
    #which consist of each key from digrah to zero num value
    indegress = {node: 0 for node in digraph}

    print(indegress)

if __name__ == "__main__":
    main()