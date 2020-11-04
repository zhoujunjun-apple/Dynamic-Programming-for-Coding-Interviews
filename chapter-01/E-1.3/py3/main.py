def move_hanoi_tower(num: int, s: str, d: str, e: str):
    """
    move num discs from s to d by the help of e. 
    larger disc cannot put upon less disc
    """
    if num == 1:
        print(f'move disc {num} from {s} to {d}')
    else:
        move_hanoi_tower(num=num-1, s=s, d=e, e=d)
        print(f'move disc {num} from {s} to {d}')
        move_hanoi_tower(num=num-1, s=e, d=d, e=s)


if __name__ == "__main__":
    num_dics = 3
    source = 's'
    destination = 'd'
    help = 'e'

    move_hanoi_tower(
        num=num_dics, 
        s=source, 
        d=destination, 
        e=help,
    )