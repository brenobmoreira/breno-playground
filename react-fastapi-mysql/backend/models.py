from sqlalchemy import Column, Integer, String
from database import Base
from pydantic import BaseModel

class Item(Base):
    __tablename__ = "items"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String(255), index=True)
    description = Column(String(255), index=True)
    price = Column(Integer)

class ItemCreate(BaseModel):
    name: str
    description: str
    price: int