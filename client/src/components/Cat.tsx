import React from 'react'
import "./Cat.css"

interface CatCardProps {
    img_url: string;
}

const CatCard: React.FC<CatCardProps> = props => {
    return (
        <div>
            <img src={props.img_url} alt="Cat" />
        </div>
    )
}

export default CatCard