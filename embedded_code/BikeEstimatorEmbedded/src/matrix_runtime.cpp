#include "matrix_runtime.hpp"

/*!
*   @brief Statically allocated thread stack. 
*/
static uint32_t matrix_thread_stack[8192]; 

/*!
*   @brief Drawing and manipulation memory. 
*/
static byte matrix_drawing_memory[NUM_STRIP_LEDS * 3]; 

/*!
*   @brief DMA memory buffer that we will use to write to the LEDs
*/
static DMAMEM byte matrix_display_memory[NUM_STRIP_LEDS * 12]; 

/*!
*   @brief WS2812b strip manipulation object. 
*/  
static WS2812Serial matrix = WS2812Serial(NUM_STRIP_LEDS, matrix_display_memory, matrix_drawing_memory, STRIP_LED_GPIO, WS2812_GRB);

// FUNCTION DECLARATIONS
void matrix_thread(void *parameters);

/*!
*   @brief Strip generation constructor. 
*/
void start_led_strip_runtime(void){
    os_add_thread(&matrix_thread, NULL, sizeof(matrix_thread_stack), matrix_thread_stack); 
}

/*!
    @brief Allows us to set pixels by matrix position
    @param uint8_t x position
    @param uint8_t y position
    @param uint8_t r color
    @param uint8_t g color
    @param uint8_t b color
*/
static void set_matrix_color(uint8_t x, uint8_t y, uint8_t r, uint8_t g, uint8_t b){
    // Bounds correction
    if(x > 15 || y > 7){
        return; 
    }

    int pos = 8 * x + y; 
    matrix.setPixelColor(pos, r, g, b);
}

/*!
*   @brief Strip thread function
*   @param void* parameter pointer to whatever we want
*/
static void matrix_thread(void *parameters){
    // If there is an issue with setting up the led strip, we just sleep the thread. 
    if(!matrix.begin())
        while(1)
            os_thread_delay_s(1);        

    matrix.setBrightness(8); 

    matrix.show(); 

    for(;;){
        // Sit and wait until we have all the data needed
        while(Serial.available() < 380){
            os_thread_sleep_ms(1);
        }

        if((Serial.read() == 16) && (Serial.read() == 24) && (Serial.read() == 33) && (Serial.read() == 22)){
            for(int x = 0; x <= 15; x++){
                for(int y = 0; y <= 7; y++){
                    uint8_t r = Serial.read(); 
                    uint8_t g = Serial.read(); 
                    uint8_t b = Serial.read(); 
                    set_matrix_color(x, y, r, g, b); 
                }
            }
        }
        matrix.show(); 
        Serial.flush();
    }
}
