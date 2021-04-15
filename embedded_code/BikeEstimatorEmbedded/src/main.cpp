#include <Arduino.h>

#include "matrix_runtime.hpp"

#include "OS/OSThreadKernel.h"

void setup() {
  os_init(); 
  Serial.begin(115200);
  start_led_strip_runtime(); 
  _os_yield(); 
}

void loop() {}