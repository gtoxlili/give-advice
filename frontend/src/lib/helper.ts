// eslint-disable-next-line @typescript-eslint/no-empty-function
export function noop() {
}

// 给颜色加上透明度
export function colorOpacity(color: string, opacity: number) {
    if (color.startsWith('#')) {
        // 如果是 3 位
        if (color.length === 4) {
            return `#${color[1]}${color[1]}${color[2]}${color[2]}${color[3]}${color[3]}${Math.round(opacity * 255).toString(16)}`;
        } else if (color.length === 7) {
            return `${color}${Math.round(opacity * 255).toString(16)}`;
        } else if (color.length === 9) {
            return `${color.substring(0, 7)}${Math.round(opacity * 255).toString(16)}`;
        }
    } else if (color.startsWith('rgb(')) {
        return `${color.replace(')', `, ${opacity})`).replace('(', 'a(')}`;
    }
    throw new Error(`不支持的颜色格式: ${color}`);
}

