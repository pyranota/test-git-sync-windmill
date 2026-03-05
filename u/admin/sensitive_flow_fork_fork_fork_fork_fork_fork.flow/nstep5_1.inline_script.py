from f.rel.branch import main as br;
from f.rel.leaf_1 import main as lf_1;
from ..leaf_1 import main as lf_12;
from ...rel.leaf_2 import main as lf_2;

def check():
    return [br(), lf_1(), lf_2(), lf_12()];

def main(x: str):
    return x