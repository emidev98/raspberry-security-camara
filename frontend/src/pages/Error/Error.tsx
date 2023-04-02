import React from 'react'

type ErrorProps = { msg: string };

export const Error: React.FC<ErrorProps> = (props) => {

    return (
        <div>{props.msg}</div>
    )
}
