# Trigrams

In natural language processing, a [_trigram_](https://en.wikipedia.org/wiki/Trigram) is a sequence of three consecutive words in a given body of text. For example, the sentence "To be or not to be, that is the question" contains the following trigrams:

```
[to, be, or]
[be, or, not]
[or, not, to]
[not, to, be]
[to, be, that]
[be, that, is]
[that, is, the]
[is, the, question]
```

Given a series of trigrams, it is possible to generate a random piece of text that resembles the original. For example, if we start with the two words "to be", we can find two trigrams that match that prefix:

```
[to, be, or]
[to, be, that]
```

At this point, we can make a random choice, and end up with "to be that". We can repeat the process by taking the last two words, "be that", and looking up what words may come next. Although with such a small example we are very limited in our choice, given a large enough body of text, such as a novel, we can produce text that, while completely devoid of meaning, appears to be in the same style as the novel that it was trained on.

For example, after "learning" the text of some of Jane Austen’s novels, this technique produced the following paragraph:

> Middleton wished it very much inclined to ask them what they did not
> know how to take her from a long drive this morning; perhaps we may
> have as many sensations of exquisite felicity followed, and the walk
> must arise from something she said herself; and Jane’s offences rose
> again. They gave themselves up wholly to their satisfaction.

## The task

We would like you to build a Go program that will "learn" the contents of one or more text files, and generate random text using trigrams as described above.

You are free to structure this program as you prefer; however, we recommend a very simple CLI so that running your program with a given text file will output a set number of words based on the trigrams present in the text file.

For example, running your program with the full text of the novel _Pride and Prejudice_ (available in the public domain at [Project Gutenberg](https://www.gutenberg.org)) should generate some random text in that style as follows:

```
$ ./trigrams pride-prejudice.txt
To think it more than commonly anxious to get round to the preference of one, and offended by the other as politely and more cheerfully. Their visit afforded was produced by the lady with whom she almost looked up to the stables. They were to set out with such a woman.
```

## Requirements

In addition to the above, we would like you to keep the following requirements in mind:

  1. The random choice of words should be proportional to the frequency of that word in the trigrams. For example, if the word "stormy" follows the words "dark and" 9 times out of 10, then we would expect the randomly-generated text to preserve that frequency.
  2. A key property of text that has been randomly generated based on a series of trigrams is that for any three consecutive words in the generated text, there should be at least one occurrence of those three words (i.e. that trigram) in the original text. If your program doesn’t uphold that property, then the text generation logic is very likely to contain a bug.
  3. When generating text, you will need to choose a random trigram to start with. You’re welcome to spend extra time on this to produce more realistic-sounding text, but it’s not a requirement of the task.
  4. For the purposes of this task, you don’t need to worry about punctuation, or the correct capitalisation of words.

## What we’d like to see

Although this is a relatively small task, we expect production-ready code that’s commented and tested as appropriate.

As well as sending us your code, we would like to hear about any design considerations or technical details you gave particular consideration to. If you had to make any particular trade-offs, or if you would have done things differently given more time, then let us know.
