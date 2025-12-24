from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal, engine
from models import Base, Item
from crud import get_item, create_item, get_all_items
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

origins = [
    "http://localhost:3000"
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"]
)

Base.metadata.create_all(bind=engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.post("/items/")
async def create_item_endpoint(name: str, description: str, price: int, db: Session = Depends(get_db)):
    return create_item(db, name, description, price)


@app.get("/items/{item_id}")
async def get_item_endpoint(item_id: int, db: Session = Depends(get_db)):
    item = get_item(db, item_id)
    if item is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return item


@app.get("/items/")
async def get_all_items_endpoint(db: Session = Depends(get_db)):
    items = get_all_items(db)
    if items is None:
        raise HTTPException(status_code=404, detail="Items not found")
    return items


@app.put("/items/{item_id}")
async def update_item_endpoint(item_id: int, name: str, description: str, price: int, db: Session = Depends(get_db)):
    item = db.query(Item).filter(Item.id == item_id).first()
    if item is None:
        raise HTTPException(status_code=404, detail="Item not found")
    
    item.name = name
    item.description = description
    item.price = price
    db.commit()
    db.refresh(item)

    return item


@app.delete("/items/{item_id}")
async def delete_item_endpoint(item_id: int, db: Session = Depends(get_db)):
    item = db.query(Item).filter(Item.id == item_id).first()
    if item is None:
        raise HTTPException(status_code=404, detail="Item not found")
    
    db.delete(item)
    db.commit()

    return {"detail": "Item deleted"}