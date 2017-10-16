module Bob exposing (..)

import String exposing ( toUpper, toLower, endsWith )

type Statement = Stating String | Asking String | Shouting String

discern : String -> Statement
discern statement =
  if endsWith "?" statement then
    Asking statement
  else if (statement == toUpper statement) && (statement /= toLower statement) then
    Shouting statement
  else
    Stating statement

hey : String -> String
hey statement =
  case discern statement of
    Asking statement -> "Sure."
    Shouting statement -> "Whoa, chill out!"
    Stating statement -> "Whatever."