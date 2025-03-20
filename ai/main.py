from fastapi import FastAPI, UploadFile, File
import aiofiles
import pdfplumber
from docx import Document
import httpx
import os

app = FastAPI()

DEEPSEEK_API_KEY = "sk-415e98bc2dfe422aa5a66b41d8c35000" 
DEEPSEEK_API_URL = "https://api.deepseek.com/summarize"

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.post("/upload/")
async def upload_cv(file: UploadFile = File(...)):
    if file.filename.endswith(".pdf"):
        text = await extract_text_from_pdf(file)
    elif file.filename.endswith(".docx"):
        text = await extract_text_from_docx(file)
    else:
        return {"error": "Unsupported file format"}
    
    summary = await summarize_text(text)
    return {"summary": summary}

async def extract_text_from_pdf(file: UploadFile):
    text = ""
    async with aiofiles.open(file.filename, 'wb') as out_file:
        content = await file.read()
        await out_file.write(content)
    
    with pdfplumber.open(file.filename) as pdf:
        for page in pdf.pages:
            text += page.extract_text() + "\n"
    
    os.remove(file.filename)
    return text

async def extract_text_from_docx(file: UploadFile):
    text = ""
    async with aiofiles.open(file.filename, 'wb') as out_file:
        content = await file.read()
        await out_file.write(content)
    
    doc = Document(file.filename)
    for para in doc.paragraphs:
        text += para.text + "\n"
    
    os.remove(file.filename)
    return text

async def summarize_text(text: str):
    async with httpx.AsyncClient() as client:
        response = await client.post(DEEPSEEK_API_URL, json={"text": text}, headers={"Authorization": f"Bearer {DEEPSEEK_API_KEY}"})
        return response.json().get("summary", "No summary available")
