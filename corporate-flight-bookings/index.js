/**
 * @param {number[][]} bookings
 * @param {number} n
 * @return {number[]}
 */

const corpFlightBookings = (bookings, n) => {
    const bookingsCount = Array(n).fill(0)
    
    for (const [i, j, k ] of bookings) {
        for (flight = i; flight <= j; flight += 1) {
            bookingsCount[flight -1] = bookingsCount[flight -1] + k;
        }
    }
    
    return bookingsCount;
};
