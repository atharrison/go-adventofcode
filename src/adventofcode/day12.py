import json

print "Day 12 Part 2"


def ProcessDay12():

    with open('data/day12input.txt', 'r') as input_file:
        data=input_file.read()

        jsonData = json.loads(data)
        total = 0

        objectsToInspect = [jsonData]

        while len(objectsToInspect) > 0:

            obj = objectsToInspect.pop()

            if isinstance(obj, list):
                print "list found"
                for next in obj :
                    objectsToInspect.append(next)
            elif isinstance(obj, dict):
                print "dict found"
                more_objs = handleDict(obj)
                for o in more_objs:
                    objectsToInspect.append(o)
            elif isinstance(obj, unicode):
                print "str found"
                pass
            elif isinstance(obj, int):
                print "int found, adding"
                total += int(obj)
            else:
                print "unhandled obj: "+str(type(obj))

        print "Total: " + str(total)

def handleDict(dictObj):
    for k, v in dictObj.items():
        if v == "red":
            return []

    result = []
    for k, v in dictObj.items():
        result.append(v)

    return result

ProcessDay12()
