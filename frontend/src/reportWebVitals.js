import { getCLS, getFID, getLCP } from 'web-vitals';

const reportWebVitals = (onPerfEntry) => {
    if (onPerfEntry && onPerfEntry instanceof Function) {
        getCLS(onPerfEntry); // Cumulative Layout Shift
        getFID(onPerfEntry); // First Input Delay
        getLCP(onPerfEntry); // Largest Contentful Paint
    }
};

export default reportWebVitals;