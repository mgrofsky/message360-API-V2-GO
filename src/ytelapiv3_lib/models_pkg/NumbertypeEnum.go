/*
 * ytelapiv3_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for NumbertypeEnum enum
 */
type NumbertypeEnum int

/**
 * Value collection for NumbertypeEnum enum
 */
const (
    Numbertype_ALL NumbertypeEnum = 1 + iota
    Numbertype_VOICE
    Numbertype_SMS
)

func (r NumbertypeEnum) MarshalJSON() ([]byte, error) { 
    s := NumbertypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *NumbertypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NumbertypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NumbertypeEnum to its string representation
 */
func NumbertypeEnumToValue(numbertypeEnum NumbertypeEnum) string {
    switch numbertypeEnum {
        case Numbertype_ALL:
    		return "all"		
        case Numbertype_VOICE:
    		return "voice"		
        case Numbertype_SMS:
    		return "sms"		
        default:
        	return "all"
    }
}

/**
 * Converts NumbertypeEnum Array to its string Array representation
*/
func NumbertypeEnumArrayToValue(numbertypeEnum []NumbertypeEnum) []string {
    convArray := make([]string,len( numbertypeEnum))
    for i:=0; i<len(numbertypeEnum);i++ {
        convArray[i] = NumbertypeEnumToValue(numbertypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NumbertypeEnumFromValue(value string) NumbertypeEnum {
    switch value {
        case "all":
            return Numbertype_ALL
        case "voice":
            return Numbertype_VOICE
        case "sms":
            return Numbertype_SMS
        default:
            return Numbertype_ALL
    }
}