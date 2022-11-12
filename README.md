# Social-Sentiment-Analyzer

It first accepts user input for the file name of which its message should be analyzed at which point it then checks a hashmap containing
key value pairs of words and their associated sentiment scores (a publicly available dataset from a study performed by Stanford University see more here: https://nlp.stanford.edu/projects/socialsent/)
and then tabulates a net sentiment score for the message by summing the sentiment scores for each word in the message that exists in the hashmap. 
Then based off of the score a star rating is given from 1-5, where 1 indicates a very negative message and 5 indicates a very positive message. 
