
from pprint import pprint
def fullJustify(words, maxWidth):
    result = []
    _result = {
        'text': [],
        'count': 0
    }
    f_text = words.pop(0)
    _result['text'].append(f_text)
    _result['count'] = len(f_text)
    for word in words:
        word_len = len(word) + 1
        if word_len + _result['count'] <= maxWidth:
            _result['text'].append(word)
            _result['count'] = _result['count'] + word_len
        else:
            if len(_result['text']) > 1:
                a = (maxWidth - _result['count'])
                b = (len(_result['text'])-1)
                sp = int(a / b) + (a % b > 0)
            else:
                sp = maxWidth - _result['count']
            _result['sp'] = sp
            _result['spa'] = maxWidth - _result['count']
            result.append(_result)
            _result = {
                'text': [],
                'count': 0
            }
            _result['text'].append(word)
            _result['count'] = len(word)

    if len(_result['text']) > 1:
        a = (maxWidth - _result['count'])
        b = (len(_result['text'])-1)
        sp = int(a / b) + (a % b > 0)
    else:
        sp = maxWidth - _result['count']
    _result['sp'] = sp
    _result['spa'] = maxWidth - _result['count']
    result.append(_result)

    datas = []
    for idx, r in enumerate(result):
        pprint(r)
        data = ""
        if idx != len(result) - 1:
            for idxx, text in enumerate(r['text']):
                if r['spa'] != 0:
                    if idxx != len(r['text'])-1:
                        if r['sp'] <= r['spa']:
                            sp = r['sp']
                        else:
                            sp = r['spa']
                        data = data + text + ' ' + (' '*sp)
                        r['spa'] = r['spa'] - sp
                    else:
                        data = ' '.join(r['text'])
                        print(len(data))
                        if len(data) < maxWidth:
                            data = data +' ' * (maxWidth - len(data)+1)
                else:
                    data = data + text + ' '
            data = data[:-1]
            datas.append(data)
        else:
            data = ' '.join(r['text'])
            if len(data) < maxWidth:
                data = data +' ' * (maxWidth - len(data))
            datas.append(data)
    print(datas)
    return datas



# words = ["The","important","thing","is","not","to","stop","questioning.","Curiosity","has","its","own","reason","for","existing."]
# words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]
words = ["What","must","be","acknowledgment","shall","be"]

maxWidth = 16
# maxWidth = 17
# maxWidth = 20
test = fullJustify(words, maxWidth)
# print(test)