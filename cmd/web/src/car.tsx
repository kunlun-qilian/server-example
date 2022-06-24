import {useListCar} from './client-bff'
import React from "react";


export default  Car
function Car() {
    const {data} = useListCar({base:"/api/v1"})
    return (
        <div className="Car">
            `{data?.map(car => car.name)}`
        </div>
    );
}