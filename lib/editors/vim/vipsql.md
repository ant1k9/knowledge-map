[vipsql](https://github.com/martingms/vipsql)

*Description*: A vim-frontend for interacting with psql

*Labels*: #VimScript #Vim

*Examples*:

```vim
" Starts an async psql job, prompting for the psql arguments.
" Also opens a scratch buffer where output from psql is directed.
noremap <leader>po :VipsqlOpenSession<CR>

" Terminates psql (happens automatically if the output buffer is closed).
noremap <silent> <leader>pk :VipsqlCloseSession<CR>

" In normal-mode, prompts for input to psql directly.
nnoremap <leader>ps :VipsqlShell<CR>

" In visual-mode, sends the selected text to psql.
vnoremap <leader>ps :VipsqlSendSelection<CR>

" Sends the selected _range_ to psql.
noremap <leader>pr :VipsqlSendRange<CR>

" Sends the current line to psql.
noremap <leader>pl :VipsqlSendCurrentLine<CR>

" Sends the entire current buffer to psql.
noremap <leader>pb :VipsqlSendBuffer<CR>

" Sends SIGINT (C-c) to the psql process.
noremap <leader>pc :VipsqlSendInterrupt<CR>
```
