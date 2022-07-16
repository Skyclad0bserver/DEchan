Function InitializePrivate() Uint64
    //10 IF SIGNER == address_raw("MYADDRESS") THEN GOTO 30
    //20 RETURN 0
    30 STORE("owner", SIGNER())
    40 STORE("IsBoard", "True")
    50 RETURN 0
End Function

Function Post(message String, reply String) Uint64
    10 dim key as String
    15 dim current_time as Uint64
    20 LET current_time = BLOCK_TIMESTAMP()
    30 LET key = "" + current_time + ":r" + reply
    40 STORE(key, message)
    50 RETURN 0
End Function
