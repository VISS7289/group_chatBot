import os

corpus_name = "cornell movie-dialogs corpus"

datafile = "formatted_movie_lines.txt"
MOVIE_LINES_FIELDS = ["lineID", "characterID", "movieID", "character", "text"]
MOVIE_CONVERSATIONS_FIELDS = ["character1ID", "character2ID", "movieID", "utteranceIDs"]

model_name = 'cb_model'
attn_model = 'dot'
#attn_model = 'general'
#attn_model = 'concat'
hidden_size = 500
encoder_n_layers = 2
decoder_n_layers = 2
dropout = 0.1
batch_size = 64

checkpoint_iter = 4000

clip = 50.0
teacher_forcing_ratio = 1.0
learning_rate = 0.0001
decoder_learning_ratio = 5.0
n_iteration = 4000
print_every = 1
save_every = 500

#PATH="cb_model/cornell movie-dialogs corpus/2-2_500/4000_checkpoint.tar"

PATH="./myChat/cb_model/cornell movie-dialogs corpus/2-2_500/60000_checkpoint.tar"
