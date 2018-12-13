-module(binarysearch).
-import(lists,[nth/2]).
-export([while/1, while/2, start/0]). 

while(L) -> while(L,0). 
while([], Acc) -> Acc;

while([_|T], Acc) ->
    L = length(T),
    if
        Acc+1 < L ->
            io:fwrite("~w~n",[nth(Acc+1, T)]);
        true ->
            io:fwrite("")
   end,
   while(T,Acc+1). 
   
   start() -> 
   X = [1,3,9,7,5,4,6,2,8], 
   while(X).