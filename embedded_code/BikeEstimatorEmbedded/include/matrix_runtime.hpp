#ifndef _MATRIX_RUNTIME_H
#define _MATRIX_RUNTIME_H

#include <Arduino.h>
#include "OS/OSThreadKernel.h"
#include "WS2812Serial.h"

/*!
*   @brief How many LEDs does the project have. 
*/
const int NUM_STRIP_LEDS = 128;  

/*!
*   @brief What gpio pad is our LED strip on. 
*/
const int STRIP_LED_GPIO = 20; 


// EXTERNAL FUNCTION DECLARATIONS
void start_led_strip_runtime(void); 

#endif 