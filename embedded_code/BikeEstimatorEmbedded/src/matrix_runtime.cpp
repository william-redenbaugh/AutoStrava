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
    if(x > 7 || y > 15){
        return; 
    }

    int pos = 8 * y + x; 
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

    set_matrix_color(0, 1, 100, 100, 100);
    
    matrix.show(); 

    for(;;){
        os_thread_sleep_ms(100);
    }
}
