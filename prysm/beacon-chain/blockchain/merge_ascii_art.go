package blockchain

var mergeAsciiArt = `                                                                                                                                                  	                                                                                                                                                      
		                                                                                                                                                      
		                                 +?$$$$$$?*;                         ;*?$$$$?*;                         +!$$$$$$?!;                                   
		                                 !##$???$@##$+                      !&#@$??$&#@*                      +@#&$????$##+                                   
		                                 !##;     +@#&;                    !##*     ;$##*                     @#$      ;&#+                                   
		                                 !##;      !##+                   ;##$        @#@                     $#&*      ++;                                   
		                                 !##;     ;@#&;                   *##*        ?##;                    ;$##&$!+;                                       
		                                 !##?!!!?$##$+                    !##+        !##+                      ;!$@##&$!;                                    
		                                 !##@@@@$$!+                      *##*        ?##;                          ;*?@#&!                                   
		                                 !##;                             ;##$        @#@                    ;?$;       ?##+                                  
		                                 !##;                              ?##!     ;?##+                    ;##+      ;$#&;                                  
		                                 !##;                               !&#&$??$&#@*                     ;&#&$$??$$&#@+                                   
		                                 +??;                                ;*?$$$$?+                        ;+!?$$$$$!+                                     
		                                                                                                ;;;;                                                  
		                                      ;+!??$$$?!*+;                                       ;*?$@&&&&&@@$!*;                                            
		                                   *?@############&$?+                                 ;!@###############&$!;                                         
		                                ;!@####&@$????$$@#####@!                             ;?&####$?*++;++*!?@####&?;                                       
		                               *@###&$*;         ;*$&###@*                          *&###@!;            ;!@###&!                                      
		                              !###&!;               ;?&###?                        *####!                  *&###?                                     
		                             !###@+                   ;$###$                      +###&+                    ;$###?                                    
		                            ;###&;                      $###?   ;;+*!??$$$$$$$$??!$###*                      ;@###*                                   
		                            !###!                        &###?$@&#####################@$?!+;                  +###$                                   
		                            $###+                     ;*?&#################################&$?*;               &##&;                                  
		                            $###+                  ;!$&########################################&$!;            &###;                                  
		                            *###?               ;!$################################################$*;        +###@                                   
		                            ;@###+            +$&####################################################&?;      $###!                                   
		                             +###&+         *$##########################################################$+  ;$###$;                                   
		                              *&###?;     +$##############################################################?*@###$;                                    
		                               +$###&?+ ;$#####################################################################?;                                     
		                                 *@####@&#####################################################################*                                       
		                                   +$&##################@?!*++*!$&###################&$?*++*!?$###############&*                                      
		                                     $###############&?+         ;!@###############@!;         ;!@##############!                                     
		                                   ;$##############&!;              *&###########&!;              !&#############!                                    
		                                   $##############@+    +@*          ;$#########$;          +@*    ;$#############!                                   
		                                  ?##############$;    *###*           $#######$;          +&##!     ?#############+                                  
		                                 +##############$     !#####!           $#####@;          *#####?     ?############@;                                 
		                                 @#############@;    !#######!          ;&####+          *#######?     $############?                                 
		                                +##############+    ?#########?          $###$          !#########$    +############&;                                
		                                $#############$   ;$###########?         !###?         !###########$;   $############!                                
		                                @#############*   !#############!        ?###$        *#############?   +############@                                
		                               ;&############&;    +?@#######&$+;       ;&####;       ;+?@#######&?+;    @############;                               
		                               ;#############@        +$&#&$*;         ;$#####@;          +$&#&$*;       $############+                               
		                               ;#############$    *+    ;+;   ;*;     *&#######&!     ;*;   ;+;    +*;   $############*                               
		                                &############@    ;$@!;     +$@!    ;?###########$;    *@$*     ;*$@+    $############*                               
		                                $############&;    ;$#&$*+!@##*    +@#############@+    +&#@?++$&#@;     @############+                               
		                                *#############*      $######&+    !#################?    +@######$;     +#############;                               
		                                 @############$       ?####@+   ;$###################$;   ;@####$       $############@                                
		                                 *#############*       !##@;   +@#####################&*   ;$##?       +#############!                                
		                                  $#############+       *$;   ?#########################?;   $!       +&############&;                                
		                                  ;&#############!          +$###########################@+          *&#############?                                 
		                                   +##############$*;     ;?###############################$+      *$##############@;                                 
		                                    *###############&$?!!$###################################$?!?$&################+                                  
		                                     *###################################&@$$$$@&#################################!                                   
		                                      +&##############################&?+;      ;+?&#############################?                                    
		                                       ;$############################@;            ;@###########################!                                     
		                                         ?###########################*              *##########################*                                      
		                                          +@#############&$!+$#######?              ?########$+!$&###########@+                                       
		                                            !&###########&;   $#######$+          +$########?   +&##########?;                                        
		                                             ;?###########&*   *@#######@$!;   ;$@########@*   *##########$+                                          
		                                               ;?&##########?;  ;*$&####&$*    ;!$&####&$*   ;$#########@*                                            
		                                                 ;!@#########@!;   ;++*+;   ;*;   ;+*++;   ;!&########$*                                              
		                                                    *$&########&$*;      ;*$&#&$*;      +*$&#######&?+                                                
		                                                      ;*$&#########&@$$@&#########&@$$@&########&$*;                                                  
		                                                         ;+?$&##############################&$?+;                                                     
		                                                             ;+!?$@&###################&$$!+;                                                         
		                                                                   ;++*!??$$$$$$$?!!*+;                                                               
		                                                                                                                                                      
		                  ;;;           ;;+*++;   ;;;++;;;++;;;  ;+++;;;++++; ;;;         ;;      ;;;      ;;;++;;;+++;;   ;;;+++++++;   ;+++++;              
		                 ;@#&+        +$&&@$@&#?  @#@@@&#&@@@#$ ;$@@@@#&@@@@! !#@        !#$     +&#&;     !#&@@@#&@@@&&;  !#&@@@@@@@*   $#&@@@&&$+           
		                 $#?#@;      ?#&!;   *#$  &&;;;$#!;;*#$   ;;;*#@;;;;   $#?      +#&;    ;@#?#$     !#!;;*#@;;;@#;  ?#$;;;;;;;    @#*   ;!&#?          
		                *#$ ?#$     *#&;     ;+;  ++   $#!  ;+;      *#$       ;&#+     $#*     ?#? $#!    ;+;  +#@   ++   ?#$           $#*     ;&#*         
		               ;&&;  @#*    $#$                $#!           *#$        !#@;   !#$     +#@  ;&#;        +#@        ?#@??????!    $#*      $#$         
		               $#!;;;*#&;   $#?                $#!           *#$         $#?  ;&&;    ;@#*;;;!#$        +#@        ?#@??????!    $#*      $#$         
		              !##@@@@@&#$   !#@;               $#!           *#$         ;&#+ $#*     ?#&@@@@@##!       +#@        ?#$           $#*      @#*         
		             ;&&+;;;;;;@#*  ;$#$+     @$       $#!           *#$          *#@!#?     *#@;;;;;;+##+      +#@        ?#$           $#*    +$#$          
		             $#*       +#&;  ;?&#@$$$$#$    +$$&#@$$;   ;$$$$@##$$$$*      $##@;    ;&#+       !#@;   $$@##$$*     ?#&$$$$$$$!   @#@$$$@&@!           
		            ;**         +*;     +*!?!*+;    ;*******    ;***********+      ;**;     ;*+         **;   *******+     ;*********+   +*!!!!*;             
		                                                                                                                                                      	                                                                                                                                                      
`
