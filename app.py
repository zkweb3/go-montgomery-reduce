import gradio as gr

def montgomery(x, y, m):
    return "montgomery"

inputs=[gr.Textbox(type="text", label="number1"), gr.Textbox(type="text", label="number2"), gr.Textbox(type="text", label="modular")]

demo = gr.Interface(fn=montgomery, inputs=inputs, outputs="text")

if __name__ == "__main__":
    demo.launch(share=False, server_name="0.0.0.0", server_port=7860)