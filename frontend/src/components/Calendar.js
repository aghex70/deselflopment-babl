import React, { useState } from 'react';

const Calendar = () => {
    const [currentDate, setCurrentDate] = useState(new Date());

    const renderTable = () => {
        const month = currentDate.getMonth();
        const year = currentDate.getFullYear();

        // Get the number of days in the current month
        const numDays = new Date(year, month + 1, 0).getDate();

        // Get the first day of the month
        const firstDay = new Date(year, month, 1).getDay();

        // Calculate the starting day of the week (Sunday = 0, Monday = 1, etc.)
        const startDay = (firstDay + 6) % 7;

        // Create an array to hold the table rows
        const rows = [];

        // Generate the table rows
        let dayCount = 1;
        for (let i = 0; i < 6; i++) {
            const cells = [];
            for (let j = 0; j < 7; j++) { // <-- Define the variable 'j' here
                const dayIndex = i * 7 + j;
                const day = dayIndex - startDay + 1;
                if (day <= 0 || day > numDays) {
                    // Empty cell before the first day of the month or after the last day
                    cells.push(
                        <td
                            key={`${i}-${j}`}
                            style={{
                                border: '1px solid black',
                                width: '14.28%',
                                height: '14.28%',
                            }}
                        ></td>
                    );
                } else {
                    // Cell with the current day
                    cells.push(
                        <td
                            key={`${i}-${j}`}
                            style={{
                                border: '1px solid black',
                                width: '14.28%',
                                height: '14.28%',
                            }}
                        >
                            <p style={{ textAlign: "center", marginTop: "-45%" }}>{day}</p>
                        </td>
                    );
                    dayCount++;
                }
            }

            const isWeekEmpty = cells.every(cell => !cell.props.children);

            if (!isWeekEmpty) {
                rows.push(<tr key={i}>{cells}</tr>);
            } else {
                // Break the loop if the whole week is empty
                break;
            }
        }

        return (
            <table style={{ borderCollapse: 'collapse', width: '100%', height: '100%' }}>
                <thead>
                <tr>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        <p style={{textAlign: "center"}} >Mon </p>
                    </th>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        Tue
                    </th>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        Wed
                    </th>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        Thu
                    </th>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        Fri
                    </th>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        Sat
                    </th>
                    <th
                        style={{
                            // border: '1px solid black',
                            width: '14.28%',
                            height: '10%',
                        }}
                    >
                        Sun
                    </th>
                </tr>
                </thead>
                <tbody>{rows}</tbody>
            </table>
        );
    };

    return (
        <div style={{ height: '80vh', width: '100vh' }}>
            {renderTable()}
        </div>
    );
};

export default Calendar;
